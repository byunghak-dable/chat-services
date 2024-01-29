package org.wid.userservice.service;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.wid.userservice.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.entity.entity.User.LoginType;
import org.wid.userservice.port.primary.AuthServicePort;
import org.wid.userservice.service.oauth2.Oauth2Service;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Mono;

@Service
@Slf4j
public class AuthService implements AuthServicePort {

  private final Map<LoginType, Oauth2Service> oauth2ServiceMap;

  public AuthService(
      @Qualifier("GoogleOauth2Service") Oauth2Service googleOauth2Service,
      @Qualifier("GithubOauth2Service") Oauth2Service githubOauth2Service) {
    oauth2ServiceMap = Map.of(
        LoginType.GOOGLE, googleOauth2Service,
        LoginType.GITHUB, githubOauth2Service);
  }

  @Override
  public Mono<Object> oauth2Login(Oauth2LoginRequestDto loginDto) {
    Oauth2Service oauth2Service = oauth2ServiceMap.get(loginDto.getType());

    return oauth2Service
        .getToken(loginDto.getCode())
        .flatMap(tokenResponseDto -> {
          log.info("accessToken: {}", tokenResponseDto.accessToken());
          return oauth2Service.getResource(tokenResponseDto.accessToken());
        });
  }
}
