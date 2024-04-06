package org.wid.userservice.adapter.driven.persistence.entity;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.wid.userservice.domain.entity.User.LoginType;

import lombok.Builder;
import lombok.Getter;

@Document(collection = "user")
@Getter
@Builder
public class UserEntity {

  @Id private final String id;

  private final String email;

  private final String name;

  private final String profile;

  private final LoginType loginType;
}
