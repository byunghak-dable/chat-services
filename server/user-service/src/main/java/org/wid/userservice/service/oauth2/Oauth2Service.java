package org.wid.userservice.service.oauth2;

import org.springframework.web.reactive.function.client.ClientResponse;
import org.wid.userservice.dto.oauth2.token.TokenResponseDto;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.exception.BadRequestException;

import reactor.core.publisher.Mono;

enum RequestType {
  TOKEN, RESOURCE
}

public interface Oauth2Service {

  Mono<TokenResponseDto> getToken(String code);

  Mono<UserDto> getResource(String accessToken);

  default Mono<? extends Throwable> handleClientErrorResponse(ClientResponse errorResponse) {
    return errorResponse.bodyToMono(String.class).map(BadRequestException::new);
  }
}
