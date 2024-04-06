package org.wid.userservice.adapter.driven.oauth2;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.adapter.driven.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.adapter.driven.oauth2.dto.TokenResponseDto;
import org.wid.userservice.adapter.driven.oauth2.dto.github.GithubTokenRequestDto;
import org.wid.userservice.adapter.driven.oauth2.dto.github.GithubUserDto;
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
  public Mono<TokenResponseDto> getToken(String code) {
    GithubTokenRequestDto requestDto =
        new GithubTokenRequestDto(
            githubProperties.getClientId(),
            githubProperties.getClientSecret(),
            githubProperties.getRedirectUri(),
            code);

    return tokenWebClient
        .post()
        .bodyValue(requestDto)
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(TokenResponseDto.class);
  }

  @Override
  public Mono<User> getResource(TokenResponseDto tokenResponseDto) {
    return resourceWebClient
        .get()
        .headers(headers -> headers.setBearerAuth(tokenResponseDto.accessToken()))
        .retrieve()
        .onStatus(HttpStatusCode::is4xxClientError, this::handleClientErrorResponse)
        .bodyToMono(GithubUserDto.class)
        .map(this::toUser);
  }

  private User toUser(GithubUserDto githubDto) {
    return new User(null, githubDto.email(), githubDto.name(), null, User.LoginType.GITHUB);
  }
}
