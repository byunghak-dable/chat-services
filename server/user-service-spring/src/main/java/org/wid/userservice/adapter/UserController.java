package org.wid.userservice.adapter;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.service.UserService;

@RestController
@RequestMapping("/api/v1")
public class UserController {
  public UserService userService;

  public UserController(UserService userService) {
    this.userService = userService;
  }

  @PostMapping("/user")
  public void uploadUser() {
    System.out.println("upload user");
  }

  @GetMapping("/user")
  public String getUser() {
    return "user";
  }
}
