package org.wid.userservice.adapter.driven.client.oauth2;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.adapter.driven.client.oauth2.dto.Oauth2TokenDto;
import org.wid.userservice.adapter.driven.client.oauth2.dto.google.GoogleTokenRequestDto;
import org.wid.userservice.adapter.driven.client.oauth2.dto.google.GoogleUserDto;
import org.wid.userservice.adapter.driven.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.port.driven.Oauth2ClientPort;
import reactor.core.publisher.Mono;

@Component
@Qualifier("GoogleOauth2Client")
public class GoogleOauth2Client implements Oauth2ClientPort {
  private final OAuth2ClientProperties googleProperties;
  private final WebClient tokenWebClient;
  private final WebClient resourceWebClient;

  public GoogleOauth2Client(OAuth2ClientProperties googleProperties) {
    this.googleProperties = googleProperties;
    this.tokenWebClient =
        WebClient.builder()
            .baseUrl(googleProperties.getTokenUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build();
    this.resourceWebClient =
        WebClient.builder()
            .baseUrl(googleProperties.getResourceUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build();
  }

  @Override
  public Mono<UserDto> getUserResource(String code) {
    return this.getToken(code).flatMap(this::getResource);
  }

  private Mono<Oauth2TokenDto> getToken(String code) {
    GoogleTokenRequestDto requestDto =
        new GoogleTokenRequestDto(
            googleProperties.getClientId(),
            googleProperties.getClientSecret(),
            googleProperties.getRedirectUri(),
            code);

    return tokenWebClient
        .post()
        .bodyValue(requestDto)
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(Oauth2TokenDto.class);
  }

  private Mono<UserDto> getResource(Oauth2TokenDto token) {
    return resourceWebClient
        .get()
        .uri(uriBuilder -> uriBuilder.queryParam("access_token", token.accessToken()).build())
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(GoogleUserDto.class)
        .map(this::toUser);
  }

  private UserDto toUser(GoogleUserDto googleUser) {
    return new UserDto(null, googleUser.email(), googleUser.name(), null, User.LoginType.GOOGLE);
  }
}
