package org.wid.userservice.adapter.driving.rest.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.application.dto.auth.AuthenticationDto;
import org.wid.userservice.application.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.port.driving.AuthServicePort;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/auth/v1")
@RequiredArgsConstructor
public class AuthController {
  private final AuthServicePort authService;

  @PostMapping("/login/oauth2")
  public Mono<AuthenticationDto> oauth2Login(
      @Valid @RequestBody Oauth2LoginRequestDto loginRequestDto) throws IllegalArgumentException {

    return authService.oauth2Login(loginRequestDto);
  }

  @GetMapping("/token/access/:refreshToken")
  public String refreshToken(String refreshToken) {
    return authService.generateAccessToken(refreshToken);
  }
}
