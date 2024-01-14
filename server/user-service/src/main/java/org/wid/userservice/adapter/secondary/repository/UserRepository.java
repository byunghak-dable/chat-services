package org.wid.userservice.adapter.secondary.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import org.wid.userservice.entity.entity.User;
import org.wid.userservice.port.secondary.UserRepositoryPort;

import lombok.RequiredArgsConstructor;

interface UserDao extends JpaRepository<User, Long> {
}

@Repository
@RequiredArgsConstructor
public class UserRepository implements UserRepositoryPort {

  private final UserDao userDao;

  @Override
  public void register(User user) {
    userDao.save(user);
  }

  @Override
  public void login() {
    // TODO: Auto-generated method stub
    throw new UnsupportedOperationException("Unimplemented method 'login'");
  }

  @Override
  public User getUser(long userId) {
    return userDao.findById(userId).orElseThrow();
  }
}
