package org.wid.userservice.port.driving;

import org.wid.userservice.application.dto.user.UserDto;

import reactor.core.publisher.Mono;

public interface UserServicePort {
  Mono<UserDto> upsertUser(UserDto userDto);

  Mono<UserDto> getUser(String userId);
}
