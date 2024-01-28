package org.wid.userservice.service.oauth2;

import org.wid.userservice.dto.oauth2.GoogleTokenResponseDto;

import reactor.core.publisher.Mono;

public interface Oauth2Service {
  enum RequestType {
    TOKEN, RESOURCE
  }

  Mono<Object> requestAccessToken(String code);
}
