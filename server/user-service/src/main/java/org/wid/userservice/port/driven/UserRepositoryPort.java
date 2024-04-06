package org.wid.userservice.port.driven;

import org.wid.userservice.domain.entity.User;
import reactor.core.publisher.Mono;

public interface UserRepositoryPort {
  Mono<User> upsertUser(User user);

  Mono<User> getUserById(String userId);
}
