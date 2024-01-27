package org.wid.userservice.adapter.secondary.repository;

import org.springframework.data.mongodb.repository.ReactiveMongoRepository;
import org.springframework.stereotype.Repository;
import org.wid.userservice.entity.entity.User;
import org.wid.userservice.port.secondary.UserRepositoryPort;

import lombok.RequiredArgsConstructor;
import reactor.core.publisher.Mono;

interface UserDao extends ReactiveMongoRepository<User, String> {
}

@Repository
@RequiredArgsConstructor
public class UserRepository implements UserRepositoryPort {

  private final UserDao userDao;

  @Override
  public Mono<User> register(User user) {
    return userDao.save(user);
  }

  @Override
  public Mono<User> getUserById(String userId) {
    return userDao.findById(userId);
  }
}
