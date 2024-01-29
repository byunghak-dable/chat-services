package org.wid.userservice.service.oauth2;

import java.util.Map;

import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.dto.oauth2.token.TokenRequestDto;
import org.wid.userservice.dto.oauth2.token.TokenResponseDto;
import org.wid.userservice.dto.user.UserDto;

import reactor.core.publisher.Mono;

@Service
public class GithubOauth2Service implements Oauth2Service {
  private final OAuth2ClientProperties githubProperties;
  private final Map<RequestType, WebClient> webClientMap;

  public GithubOauth2Service(OAuth2ClientProperties githubProperties) {
    this.githubProperties = githubProperties;
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
    TokenRequestDto requestDto = new TokenRequestDto(
        githubProperties.getClientId(),
        githubProperties.getClientSecret(),
        code);

    return webClientMap.get(RequestType.TOKEN)
        .post()
        .bodyValue(requestDto)
        .retrieve()
        .bodyToMono(TokenResponseDto.class);
  }

  @Override
  public Mono<UserDto> getResource(String accessToken) {
    return webClientMap.get(RequestType.RESOURCE)
        .get()
        .headers(headers -> headers.setBearerAuth(accessToken))
        .retrieve()
        .bodyToMono(UserDto.class);
  }
}
