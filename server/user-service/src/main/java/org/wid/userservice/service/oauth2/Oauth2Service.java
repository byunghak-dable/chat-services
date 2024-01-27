package org.wid.userservice.service.oauth2;

import org.wid.userservice.dto.user.OauthLoginResponseDto;

import reactor.core.publisher.Mono;

public interface Oauth2Service {
  Mono<OauthLoginResponseDto> getToken(String code);
}
