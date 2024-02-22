package org.wid.userservice.port.driving;

import org.wid.userservice.application.dto.auth.AccessTokenDto;
import org.wid.userservice.application.dto.auth.AuthenticationTokensDto;
import org.wid.userservice.application.dto.auth.Oauth2LoginRequestDto;

import reactor.core.publisher.Mono;

public interface AuthServicePort {
  Mono<AuthenticationTokensDto> oauth2Login(Oauth2LoginRequestDto dto);

  AccessTokenDto generateAccessToken(String refreshToken);
}
