package org.wid.userservice.adapter.primary.rest.controller;

import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.dto.user.OauthLoginResponseDto;
import org.wid.userservice.port.primary.AuthServicePort;

import lombok.RequiredArgsConstructor;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/auth/v1")
@RequiredArgsConstructor
public class AuthController {
  private final AuthServicePort authService;

  @PostMapping("/login/google/code/{code}")
  public Mono<OauthLoginResponseDto> googleLogin(@PathVariable String code) {
    return authService.googleLogin(code);
  }

  @PostMapping("/login/github/code/{code}")
  public Mono<OauthLoginResponseDto> githubLogin(@PathVariable String code) {
    return authService.githubLogin(code);
  }
}
