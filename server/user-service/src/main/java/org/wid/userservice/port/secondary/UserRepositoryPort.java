package org.wid.userservice.port.secondary;

import org.wid.userservice.entity.entity.User;

public interface UserRepositoryPort {
  void register(User user);

  void login();

  User getUser(long userId);
}
