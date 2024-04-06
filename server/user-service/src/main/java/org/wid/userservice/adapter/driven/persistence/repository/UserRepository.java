package org.wid.userservice.adapter.driven.persistence.repository;

import lombok.RequiredArgsConstructor;
import org.springframework.data.mongodb.repository.ReactiveMongoRepository;
import org.springframework.stereotype.Repository;
import org.wid.userservice.adapter.driven.persistence.orm.UserEntity;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.port.driven.UserRepositoryPort;
import reactor.core.publisher.Mono;

interface UserDao extends ReactiveMongoRepository<UserEntity, String> {}

@Repository
@RequiredArgsConstructor
public class UserRepository implements UserRepositoryPort {

  private final UserDao userDao;

  @Override
  public Mono<User> upsertUser(User user) {
    return userDao.save(UserEntity.from(user)).map(UserEntity::toDomain);
  }

  @Override
  public Mono<User> getUserById(String userId) {
    return userDao.findById(userId).map(UserEntity::toDomain);
  }
}
