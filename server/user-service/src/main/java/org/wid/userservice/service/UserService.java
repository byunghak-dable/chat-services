package org.wid.userservice.service;

import org.springframework.stereotype.Service;
import org.wid.userservice.dto.user.RegisterUserDto;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.entity.entity.User;
import org.wid.userservice.mapper.UserMapper;
import org.wid.userservice.port.primary.UserServicePort;
import org.wid.userservice.port.secondary.UserRepositoryPort;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class UserService implements UserServicePort {

  private final UserRepositoryPort userRepository;
  private final UserMapper userMapper;

  @Override
  public void register(RegisterUserDto registerUserDto) {
    User user = userMapper.toEntity(registerUserDto);

    userRepository.register(user);
  }

  @Override
  public void login() {
    throw new UnsupportedOperationException("Unimplemented method 'signin'");
  }

  @Override
  public UserDto getUser(long userId) {
    User user = userRepository.getUser(userId);

    return userMapper.fromEntity(user);

  }
}
