package org.wid.userservice.service.oauth2;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.dto.oauth2.resource.GoogleUserDto;
import org.wid.userservice.dto.oauth2.token.TokenRequestDto;
import org.wid.userservice.dto.oauth2.token.TokenResponseDto;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.mapper.UserMapper;

import reactor.core.publisher.Mono;

@Service
@Qualifier("GoogleOauth2Service")
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
    TokenRequestDto tokenRequestDto = new TokenRequestDto(
        googleProperties.getClientId(),
        googleProperties.getClientSecret(),
        code);

    return oauthClientMap.get(RequestType.TOKEN)
        .post()
        .bodyValue(tokenRequestDto)
        .retrieve()
        .bodyToMono(TokenResponseDto.class);
  }

  @Override
  public Mono<UserDto> getResource(String accessToken) {
    return oauthClientMap.get(RequestType.RESOURCE)
        .get()
        .uri(uriBuilder -> uriBuilder
            .queryParam("access_token", accessToken)
            .build())
        .retrieve()
        .bodyToMono(GoogleUserDto.class)
        .map(userMapper::googleUserDtoToUserDto);
  }
}
