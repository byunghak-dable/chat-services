package org.wid.userservice.application.abstraction;

import org.springframework.web.reactive.function.client.ClientResponse;
import org.wid.userservice.application.dto.oauth2.token.TokenResponseDto;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.application.exception.BadRequestException;

import reactor.core.publisher.Mono;

public interface Oauth2Service {
  enum RequestType {
    TOKEN, RESOURCE
  }

  Mono<TokenResponseDto> getToken(String code);

  Mono<UserDto> getResource(TokenResponseDto tokenResponseDto);

  default Mono<? extends Throwable> handleClientErrorResponse(ClientResponse errorResponse) {
    return errorResponse.bodyToMono(String.class).map(BadRequestException::new);
  }
}
