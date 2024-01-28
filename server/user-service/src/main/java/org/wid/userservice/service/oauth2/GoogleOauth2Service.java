package org.wid.userservice.service.oauth2;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.dto.oauth2.TokenRequestDto;
import org.wid.userservice.dto.oauth2.TokenResponseDto;

import reactor.core.publisher.Mono;

@Service
@Qualifier("GoogleOauth2Service")
public class GoogleOauth2Service implements Oauth2Service {

  private final OAuth2ClientProperties googleProperties;
  private final Map<RequestType, WebClient> oauthClientMap;

  public GoogleOauth2Service(OAuth2ClientProperties googleProperties) {
    this.googleProperties = googleProperties;
    this.oauthClientMap = Map.of(
        RequestType.TOKEN, WebClient.builder().baseUrl(googleProperties.getTokenUri()).build(),
        RequestType.RESOURCE, WebClient.builder().baseUrl(googleProperties.getResourceUri()).build());
  }

  @Override
  public Mono<TokenResponseDto> getToken(String code) {
    TokenRequestDto tokenRequestDto = new TokenRequestDto(
        googleProperties.getClientId(),
        googleProperties.getClientSecret(),
        code);

    return oauthClientMap.get(RequestType.TOKEN)
        .post()
        .accept(MediaType.APPLICATION_JSON)
        .bodyValue(tokenRequestDto)
        .retrieve()
        .bodyToMono(TokenResponseDto.class);
  }

  @Override
  public Mono<Object> getResource(String accessToken) {
    return oauthClientMap.get(RequestType.RESOURCE)
        .get()
        .uri(uriBuilder -> uriBuilder
            .queryParam("access_token", accessToken)
            .build())
        .accept(MediaType.APPLICATION_JSON)
        .retrieve()
        .bodyToMono(Object.class);
  }
}
