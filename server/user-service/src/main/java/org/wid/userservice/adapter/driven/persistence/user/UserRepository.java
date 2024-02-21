package org.wid.userservice.adapter.driven.persistence.user;

import org.springframework.stereotype.Repository;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.port.driven.UserRepositoryPort;

import lombok.RequiredArgsConstructor;
import reactor.core.publisher.Mono;

@Repository
@RequiredArgsConstructor
public class UserRepository implements UserRepositoryPort {

  private final UserDao userDao;

  @Override
  public Mono<User> upsertUser(User user) {
    // TODO: need to check if upserting is working in mongodb
    return userDao.save(user);
  }

  @Override
  public Mono<User> getUserById(String userId) {
    return userDao.findById(userId);
  }
}
