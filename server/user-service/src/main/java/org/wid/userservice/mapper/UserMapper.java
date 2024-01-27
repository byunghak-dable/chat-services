package org.wid.userservice.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.NullValueMappingStrategy;
import org.mapstruct.NullValuePropertyMappingStrategy;
import org.mapstruct.ReportingPolicy;
import org.wid.userservice.dto.user.RegisterUserDto;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.entity.entity.User;

import reactor.core.publisher.Mono;

@Mapper(componentModel = "spring", unmappedSourcePolicy = ReportingPolicy.IGNORE, nullValueMappingStrategy = NullValueMappingStrategy.RETURN_DEFAULT, nullValuePropertyMappingStrategy = NullValuePropertyMappingStrategy.SET_TO_DEFAULT)
public interface UserMapper {

  @Mapping(target = "id", ignore = true)
  User registerDtoToEntity(RegisterUserDto dto);

  UserDto entityToUserDto(User user);

  default Mono<User> registerDtoToEntityAsync(RegisterUserDto dto) {
    return Mono.just(registerDtoToEntity(dto));
  }

  default Mono<UserDto> entityToUserDtoAsync(Mono<User> user) {
    return user.map(this::entityToUserDto);
  }
}
