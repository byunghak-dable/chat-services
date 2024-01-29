package org.wid.userservice.port.primary;

import org.wid.userservice.dto.user.UserDto;

import reactor.core.publisher.Mono;

public interface UserServicePort {
  Mono<UserDto> getUser(String userId);
}
