package org.wid.userservice.service.oauth2;

import org.wid.userservice.dto.oauth2.token.TokenResponseDto;
import org.wid.userservice.dto.user.UserDto;

import reactor.core.publisher.Mono;

enum RequestType {
  TOKEN, RESOURCE
}

public interface Oauth2Service {

  Mono<TokenResponseDto> getToken(String code);

  Mono<UserDto> getResource(String accessToken);
}
