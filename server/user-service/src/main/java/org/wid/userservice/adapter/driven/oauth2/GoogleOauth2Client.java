package org.wid.userservice.adapter.driven.oauth2;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.adapter.driven.oauth2.dto.google.GoogleTokenRequestDto;
import org.wid.userservice.adapter.driven.oauth2.dto.google.GoogleUserDto;
import org.wid.userservice.adapter.driven.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.adapter.driven.oauth2.dto.TokenResponseDto;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.port.driven.Oauth2ClientPort;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Mono;

@Service
@Qualifier("GoogleOauth2Client")
@Slf4j
public class GoogleOauth2Client implements Oauth2ClientPort {
  private final OAuth2ClientProperties googleProperties;
  private final Map<RequestType, WebClient> oauthClientMap;

  public GoogleOauth2Client(OAuth2ClientProperties googleProperties) {
    this.googleProperties = googleProperties;
    this.oauthClientMap =
        Map.of(
            RequestType.TOKEN,
                WebClient.builder()
                    .baseUrl(googleProperties.getTokenUri())
                    .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
                    .build(),
            RequestType.RESOURCE,
                WebClient.builder()
                    .baseUrl(googleProperties.getResourceUri())
                    .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
                    .build());
  }

  @Override
  public Mono<TokenResponseDto> getToken(String code) {
    GoogleTokenRequestDto requestDto =
        new GoogleTokenRequestDto(
            googleProperties.getClientId(),
            googleProperties.getClientSecret(),
            googleProperties.getRedirectUri(),
            code);

    log.info("google token req body: {}", requestDto);
    return oauthClientMap
        .get(RequestType.TOKEN)
        .post()
        .bodyValue(requestDto)
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(TokenResponseDto.class);
  }

  @Override
  public Mono<User> getResource(TokenResponseDto tokenResponseDto) {
    return oauthClientMap
        .get(RequestType.RESOURCE)
        .get()
        .uri(
            uriBuilder ->
                uriBuilder.queryParam("access_token", tokenResponseDto.accessToken()).build())
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(GoogleUserDto.class)
        .map(this::toUser);
  }

  private User toUser(GoogleUserDto googleUserDto) {
    return new User(null, googleUserDto.email(), googleUserDto.name(), null, User.LoginType.GOOGLE);
  }
}
