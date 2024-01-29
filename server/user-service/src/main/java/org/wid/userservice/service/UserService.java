package org.wid.userservice.service;

import org.springframework.stereotype.Service;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.entity.entity.User;
import org.wid.userservice.mapper.UserMapper;
import org.wid.userservice.port.primary.UserServicePort;
import org.wid.userservice.port.secondary.UserRepositoryPort;

import lombok.RequiredArgsConstructor;
import reactor.core.publisher.Mono;

@Service
@RequiredArgsConstructor
public class UserService implements UserServicePort {

  private final UserRepositoryPort userRepository;
  private final UserMapper userMapper;

  @Override
  public Mono<UserDto> getUser(String userId) {
    Mono<User> user = userRepository.getUserById(userId);

    return user.map(userMapper::entityToUserDto);
  }
}
