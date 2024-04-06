package org.wid.userservice.port.driving;

import org.wid.userservice.application.dto.auth.AuthenticationDto;
import org.wid.userservice.application.dto.auth.Oauth2LoginRequestDto;
import reactor.core.publisher.Mono;

public interface AuthServicePort {
  Mono<AuthenticationDto> oauth2Login(Oauth2LoginRequestDto dto);

  String generateAccessToken(String refreshToken);
}
