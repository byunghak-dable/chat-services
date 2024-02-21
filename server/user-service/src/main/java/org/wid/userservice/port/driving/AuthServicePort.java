package org.wid.userservice.port.driving;

import org.wid.userservice.application.dto.auth.JwtDto;
import org.wid.userservice.application.dto.auth.Oauth2LoginRequestDto;

import reactor.core.publisher.Mono;

public interface AuthServicePort {
  Mono<JwtDto> oauth2Login(Oauth2LoginRequestDto dto);
}
