package org.wid.userservice.port.primary;

import org.wid.userservice.dto.auth.Oauth2LoginRequestDto;

import reactor.core.publisher.Mono;

public interface AuthServicePort {
  Mono<Object> oauth2Login(Oauth2LoginRequestDto dto);
}
