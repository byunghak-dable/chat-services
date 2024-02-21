package org.wid.userservice.adapter.driving.rest.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.application.dto.auth.JwtDto;
import org.wid.userservice.application.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.port.driving.AuthServicePort;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/auth/v1")
@RequiredArgsConstructor
public class AuthController {
  private final AuthServicePort authService;

  @PostMapping("/login/oauth2")
  public Mono<JwtDto> oauth2Login(@Valid @RequestBody Oauth2LoginRequestDto loginRequestDto)
      throws IllegalArgumentException {

    return authService.oauth2Login(loginRequestDto);
  }
}
