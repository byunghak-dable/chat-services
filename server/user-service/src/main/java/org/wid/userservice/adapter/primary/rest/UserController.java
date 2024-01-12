package org.wid.userservice.adapter.primary.rest;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.dto.user.RegisterUserDto;
import org.wid.userservice.port.primary.UserServicePort;

@RestController
@RequestMapping("/api/v1")
public class UserController {
  private final UserServicePort userService;

  public UserController(UserServicePort userService) {
    this.userService = userService;
  }

  @PostMapping("/user")
  public ResponseEntity<Void> uploadUser(@Validated @RequestBody RegisterUserDto registerUserDto) {
    userService.register(registerUserDto);

    return ResponseEntity.ok().build();
  }

  public ResponseEntity<String> login() {

    return ResponseEntity.status(HttpStatus.NOT_IMPLEMENTED).body("login");
  }

  @GetMapping("/user/{userId}")
  public ResponseEntity<String> getUser(@PathVariable int userId) {
    System.out.println("checking get user" + userId);

    return ResponseEntity.status(HttpStatus.NOT_IMPLEMENTED).body("getUser");
  }
}
