package org.wid.userservice.port.primary;

import org.wid.userservice.dto.user.RegisterUserDto;
import org.wid.userservice.dto.user.UserDto;

public interface UserServicePort {
  void register(RegisterUserDto registerUserDto);

  void login();

  UserDto getUser(long userId);
}
