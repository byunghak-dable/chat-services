package org.wid.userservice.port.driven;

import org.springframework.web.reactive.function.client.ClientResponse;
import org.wid.userservice.adapter.driven.oauth2.dto.TokenResponseDto;
import org.wid.userservice.application.exception.BadRequestException;
import org.wid.userservice.domain.entity.User;
import reactor.core.publisher.Mono;

public interface Oauth2ClientPort {
  enum RequestType {
    TOKEN,
    RESOURCE
  }

  Mono<TokenResponseDto> getToken(String code);

  Mono<User> getResource(TokenResponseDto tokenResponseDto);

  default Mono<? extends Throwable> handleClientErrorResponse(ClientResponse errorResponse) {
    return errorResponse.bodyToMono(String.class).map(BadRequestException::new);
  }
}
