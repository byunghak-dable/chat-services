package org.wid.userservice.port.primary;

import org.wid.userservice.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.dto.oauth2.TokenResponseDto;

import reactor.core.publisher.Mono;

public interface AuthServicePort {
  Mono<Void> oauth2Login(Oauth2LoginRequestDto dto);
}
