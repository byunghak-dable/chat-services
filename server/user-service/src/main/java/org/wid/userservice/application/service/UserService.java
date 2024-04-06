package org.wid.userservice.application.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.application.mapper.UserMapper;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.port.driven.UserRepositoryPort;
import org.wid.userservice.port.driving.UserServicePort;
import reactor.core.publisher.Mono;

@Service
@RequiredArgsConstructor
public class UserService implements UserServicePort {

  private final UserRepositoryPort userRepository;
  private final UserMapper userMapper;

  @Override
  public Mono<UserDto> upsertUser(UserDto userDto) {
    User user = userMapper.userDtoToEntity(userDto);

    return userRepository.upsertUser(user).map(userMapper::entityToUserDto);
  }

  @Override
  public Mono<UserDto> getUser(String userId) {
    Mono<User> user = userRepository.getUserById(userId);

    return user.map(userMapper::entityToUserDto);
  }
}
