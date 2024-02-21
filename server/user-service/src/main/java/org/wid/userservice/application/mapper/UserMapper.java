package org.wid.userservice.application.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.NullValueMappingStrategy;
import org.mapstruct.NullValuePropertyMappingStrategy;
import org.mapstruct.ReportingPolicy;
import org.wid.userservice.application.dto.oauth2.resource.GithubUserDto;
import org.wid.userservice.application.dto.oauth2.resource.GoogleUserDto;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.domain.entity.User;

@Mapper(componentModel = "spring", unmappedSourcePolicy = ReportingPolicy.IGNORE, nullValueMappingStrategy = NullValueMappingStrategy.RETURN_DEFAULT, nullValuePropertyMappingStrategy = NullValuePropertyMappingStrategy.SET_TO_DEFAULT)
public interface UserMapper {

  User userDtoToEntity(UserDto dto);

  UserDto entityToUserDto(User user);

  @Mapping(target = "id", ignore = true)
  @Mapping(target = "loginType", expression = "java(User.LoginType.GOOGLE)")
  @Mapping(target = "profile", source = "picture")
  UserDto googleUserDtoToUserDto(GoogleUserDto dto);

  @Mapping(target = "id", ignore = true)
  @Mapping(target = "loginType", expression = "java(User.LoginType.GITHUB)")
  @Mapping(target = "profile", ignore = true)
  UserDto githubUserDtoToUserDto(GithubUserDto dto);
}
