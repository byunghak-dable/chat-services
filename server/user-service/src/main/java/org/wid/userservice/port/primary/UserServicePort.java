package org.wid.userservice.port.primary;

import org.wid.userservice.dto.user.RegisterUserDto;

public interface UserServicePort {
  void register(RegisterUserDto registerUserDto);

  void login();
}
