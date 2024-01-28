package org.wid.userservice.service.oauth2;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import org.wid.userservice.config.Oauth2ClientConfig.OAuth2ClientProperties;
import org.wid.userservice.dto.oauth2.GoogleTokenRequestDto;
import org.wid.userservice.dto.oauth2.GoogleTokenResponseDto;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Mono;

@Service
@Qualifier("GoogleOauth2Service")
@Slf4j
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
  public Mono<Object> requestAccessToken(String code) {
    GoogleTokenRequestDto tokenRequestDto = new GoogleTokenRequestDto(
        googleProperties.getClientId(),
        googleProperties.getClientSecret(),
        code);

    return oauthClientMap.get(RequestType.TOKEN)
        .post()
        .accept(MediaType.APPLICATION_JSON)
        .bodyValue(tokenRequestDto)
        .retrieve()
        .bodyToMono(Object.class);
  }
}
