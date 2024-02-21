package org.wid.userservice.application.service.oauth2;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.adapter.driven.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.application.dto.oauth2.resource.GoogleUserDto;
import org.wid.userservice.application.dto.oauth2.token.GoogleTokenRequestDto;
import org.wid.userservice.application.dto.oauth2.token.TokenResponseDto;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.application.mapper.UserMapper;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Mono;

@Service
@Qualifier("GoogleOauth2Service")
@Slf4j
public class GoogleOauth2Service implements Oauth2Service {
  private final OAuth2ClientProperties googleProperties;
  private final UserMapper userMapper;
  private final Map<RequestType, WebClient> oauthClientMap;

  public GoogleOauth2Service(OAuth2ClientProperties googleProperties, UserMapper userMapper) {
    this.googleProperties = googleProperties;
    this.userMapper = userMapper;
    this.oauthClientMap = Map.of(
        RequestType.TOKEN, WebClient.builder()
            .baseUrl(googleProperties.getTokenUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build(),
        RequestType.RESOURCE, WebClient.builder()
            .baseUrl(googleProperties.getResourceUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build());
  }

  @Override
  public Mono<TokenResponseDto> getToken(String code) {
    GoogleTokenRequestDto requestDto = new GoogleTokenRequestDto(
        googleProperties.getClientId(),
        googleProperties.getClientSecret(),
        googleProperties.getRedirectUri(),
        code);

    log.info("google token req body: {}", requestDto);
    return oauthClientMap.get(RequestType.TOKEN)
        .post()
        .bodyValue(requestDto)
        .retrieve()
        .onStatus(status -> status.is4xxClientError(), this::handleClientErrorResponse)
        .bodyToMono(TokenResponseDto.class);
  }

  @Override
  public Mono<UserDto> getResource(TokenResponseDto tokenResponseDto) {
    return oauthClientMap.get(RequestType.RESOURCE)
        .get()
        .uri(uriBuilder -> uriBuilder
            .queryParam("access_token", tokenResponseDto.accessToken())
            .build())
        .retrieve()
        .onStatus(status -> status.is4xxClientError(), this::handleClientErrorResponse)
        .bodyToMono(GoogleUserDto.class)
        .map(userMapper::googleUserDtoToUserDto);
  }
}
