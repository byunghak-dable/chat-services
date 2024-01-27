package org.wid.userservice.port.primary;

import org.wid.userservice.dto.user.OauthLoginResponseDto;

import reactor.core.publisher.Mono;

public interface AuthServicePort {
  Mono<OauthLoginResponseDto> googleLogin(String code);

  Mono<OauthLoginResponseDto> githubLogin(String code);
}
