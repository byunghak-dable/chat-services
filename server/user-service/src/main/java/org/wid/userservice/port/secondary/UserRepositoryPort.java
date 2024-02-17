package org.wid.userservice.port.secondary;

import org.wid.userservice.entity.User;

import reactor.core.publisher.Mono;

public interface UserRepositoryPort {
  Mono<User> upsertUser(User user);

  Mono<User> getUserById(String userId);
}
