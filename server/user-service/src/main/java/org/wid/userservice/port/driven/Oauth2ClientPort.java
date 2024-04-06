package org.wid.userservice.port.driven;

import org.springframework.web.reactive.function.client.ClientResponse;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.application.exception.BadRequestException;
import reactor.core.publisher.Mono;

public interface Oauth2ClientPort {
  Mono<UserDto> getUserResource(String code);

  default Mono<? extends Throwable> handleClientErrorResponse(ClientResponse errorResponse) {
    return errorResponse.bodyToMono(String.class).map(BadRequestException::new);
  }
}
