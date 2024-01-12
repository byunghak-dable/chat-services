package org.wid.userservice.service;

import org.springframework.stereotype.Service;
import org.wid.userservice.dto.user.RegisterUserDto;
import org.wid.userservice.port.primary.UserServicePort;

@Service
public class UserService implements UserServicePort {

  @Override
  public void register(RegisterUserDto registerUserDto) {
    System.out.println("register successfully");
  }

  @Override
  public void login() {
    throw new UnsupportedOperationException("Unimplemented method 'signin'");
  }
}
