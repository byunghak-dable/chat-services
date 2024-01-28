package org.wid.userservice.service.oauth2;

import org.wid.userservice.dto.oauth2.TokenResponseDto;

import reactor.core.publisher.Mono;

public interface Oauth2Service {
  enum RequestType {
    TOKEN, RESOURCE
  }

  Mono<TokenResponseDto> requestAccessToken(String code);
}
