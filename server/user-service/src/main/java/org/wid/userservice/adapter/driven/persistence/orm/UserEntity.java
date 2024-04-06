package org.wid.userservice.adapter.driven.persistence.orm;

import lombok.Builder;
import lombok.Getter;
import lombok.RequiredArgsConstructor;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.domain.entity.User.LoginType;

@Document(collection = "user")
@Getter
@Builder
@RequiredArgsConstructor
public class UserEntity {

  @Id private final String id;

  private final String email;

  private final String name;

  private final String profile;

  private final LoginType loginType;

  public static UserEntity from(User user) {
    return new UserEntity(
        user.getId(), user.getEmail(), user.getName(), user.getProfile(), user.getLoginType());
  }

  public User toDomain() {
    return new User(id, email, name, profile, loginType);
  }
}
