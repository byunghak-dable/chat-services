package org.wid.userservice.application.service.oauth2;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.adapter.driven.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.application.abstraction.Oauth2Service;
import org.wid.userservice.application.dto.oauth2.resource.GithubUserDto;
import org.wid.userservice.application.dto.oauth2.token.GithubTokenRequestDto;
import org.wid.userservice.application.dto.oauth2.token.TokenResponseDto;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.application.mapper.UserMapper;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Mono;

@Service
@Qualifier("GithubOauth2Service")
@Slf4j
public class GithubOauth2Service implements Oauth2Service {
  private final OAuth2ClientProperties githubProperties;
  private final UserMapper userMapper;
  private final Map<RequestType, WebClient> webClientMap;

  public GithubOauth2Service(OAuth2ClientProperties githubProperties, UserMapper userMapper) {
    this.githubProperties = githubProperties;
    this.userMapper = userMapper;
    this.webClientMap = Map.of(
        RequestType.TOKEN, WebClient.builder()
            .baseUrl(githubProperties.getTokenUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build(),
        RequestType.RESOURCE, WebClient.builder()
            .baseUrl(githubProperties.getResourceUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build());
  }

  @Override
  public Mono<TokenResponseDto> getToken(String code) {
    GithubTokenRequestDto requestDto = new GithubTokenRequestDto(
        githubProperties.getClientId(),
        githubProperties.getClientSecret(),
        githubProperties.getRedirectUri(),
        code);

    log.info("github token req body: {}", requestDto);
    return webClientMap.get(RequestType.TOKEN)
        .post()
        .bodyValue(requestDto)
        .retrieve()
        .onStatus(status -> status.is4xxClientError(), this::handleClientErrorResponse)
        .bodyToMono(TokenResponseDto.class);
  }

  @Override
  public Mono<UserDto> getResource(TokenResponseDto tokenResponseDto) {
    return webClientMap.get(RequestType.RESOURCE)
        .get()
        .headers(headers -> headers.setBearerAuth(tokenResponseDto.accessToken()))
        .retrieve()
        .onStatus(status -> status.is4xxClientError(), this::handleClientErrorResponse)
        .bodyToMono(GithubUserDto.class)
        .map(userMapper::githubUserDtoToUserDto);
  }
}
