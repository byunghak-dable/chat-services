package org.wid.userservice.application.service;

import java.util.Map;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.wid.userservice.application.dto.auth.AuthenticationDto;
import org.wid.userservice.application.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.application.mapper.UserMapper;
import org.wid.userservice.domain.entity.Authentication;
import org.wid.userservice.domain.entity.User;
import org.wid.userservice.domain.entity.User.LoginType;
import org.wid.userservice.port.driven.Oauth2ClientPort;
import org.wid.userservice.port.driven.UserRepositoryPort;
import org.wid.userservice.port.driving.AuthServicePort;
import reactor.core.publisher.Mono;

@Service
public class AuthService implements AuthServicePort {
  private final Map<LoginType, Oauth2ClientPort> oauth2ServiceMap;
  private final UserRepositoryPort userRepository;
  private final UserMapper userMapper;

  public AuthService(
      @Qualifier("GoogleOauth2Client") Oauth2ClientPort googleClient,
      @Qualifier("GithubOauth2Client") Oauth2ClientPort githubClient,
      UserRepositoryPort userRepository,
      UserMapper userMapper) {
    this.oauth2ServiceMap = Map.of(LoginType.GOOGLE, googleClient, LoginType.GITHUB, githubClient);
    this.userRepository = userRepository;
    this.userMapper = userMapper;
  }

  @Override
  public Mono<AuthenticationDto> oauth2Login(Oauth2LoginRequestDto loginDto) {
    return oauth2ServiceMap
        .get(loginDto.getType())
        .getUserResource(loginDto.getCode())
        .map(userMapper::toEntity)
        .flatMap(userRepository::upsertUser)
        .map(this::generateTokens);
  }

  private AuthenticationDto generateTokens(User user) {
    return new AuthenticationDto(
        Authentication.accessToken(user.id()).toJsonWebToken(),
        Authentication.refreshToken(user.id()).toJsonWebToken());
  }

  @Override
  public String generateAccessToken(String refreshToken) {
    return Authentication.from(refreshToken).toJsonWebToken();
  }
}
