package org.wid.userservice.adapter.driven.client.oauth2;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.adapter.driven.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.adapter.driven.client.oauth2.dto.Oauth2TokenDto;
import org.wid.userservice.adapter.driven.client.oauth2.dto.github.GithubTokenRequestDto;
import org.wid.userservice.adapter.driven.client.oauth2.dto.github.GithubUserDto;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.port.driven.Oauth2ClientPort;
import reactor.core.publisher.Mono;

@Component
@Qualifier("GithubOauth2Client")
public class GithubOauth2Client implements Oauth2ClientPort {
  private final OAuth2ClientProperties githubProperties;
  private final WebClient tokenWebClient;
  private final WebClient resourceWebClient;

  public GithubOauth2Client(OAuth2ClientProperties githubProperties) {
    this.githubProperties = githubProperties;
    this.tokenWebClient =
        WebClient.builder()
            .baseUrl(githubProperties.getTokenUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build();
    this.resourceWebClient =
        WebClient.builder()
            .baseUrl(githubProperties.getResourceUri())
            .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
            .build();
  }

  @Override
  public Mono<UserDto> getUserResource(String code) {
    return this.getToken(code).flatMap(this::getResource);
  }

  private Mono<Oauth2TokenDto> getToken(String code) {
    GithubTokenRequestDto tokenRequest =
        new GithubTokenRequestDto(
            githubProperties.getClientId(),
            githubProperties.getClientSecret(),
            githubProperties.getRedirectUri(),
            code);

    return tokenWebClient
        .post()
        .bodyValue(tokenRequest)
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(Oauth2TokenDto.class);
  }

  private Mono<UserDto> getResource(Oauth2TokenDto token) {
    return resourceWebClient
        .get()
        .headers(headers -> headers.setBearerAuth(token.accessToken()))
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(GithubUserDto.class)
        .map(this::toUser);
  }

  private UserDto toUser(GithubUserDto githubUser) {
    return new UserDto(null, githubUser.email(), githubUser.name(), null, User.LoginType.GITHUB);
  }
}
