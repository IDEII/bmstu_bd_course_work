import { Component, OnInit, ViewChild } from '@angular/core';
import { UserService } from '../user.service';
import { Router } from '@angular/router';
import { User } from '../user';
import { RegistrationComponent } from '../registration/registration.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  // @ViewChild(RegistrationComponent) registrationComponent!: RegistrationComponent;
  title = 'app';
  id: string = "";
  user: User = {
    user_profile_id: '',
    username: '',
    profile_type: 0,
    e_mail: '',
    linked_id_1: "",
    linked_id_2: "",
    linked_id_3: "",
    password: ''
  };
  is_visible = false;
  is_visible2 = false;

  _isLoggedIn = false;
  showLogin: boolean = false;
  showRegister: boolean = false;

  constructor(private userService: UserService, private router: Router) {}

  async ngOnInit() {
    this.updateLoginStatus();
    if (this._isLoggedIn) {
      await this.loadUser(this.id);
    }
    this.checkLoginStatus();
  }

  async loadUser(id: string) {
    try {
      const data = await this.userService.getUser(id).toPromise();
      this.user = data;
      console.log('Данные пользователя загружены:', this.user);
    } catch (error) {
      console.error('Ошибка при загрузке пользователя:', error);
    }
  }

  onRegistrationComplete() {
    this.is_visible = !this.is_visible;
    this.checkLoginStatus();
  }
  onLoginComplete() {
    this.is_visible2 = !this.is_visible2; 
    this.checkLoginStatus();
  }

  checkLoginStatus() {
    this._isLoggedIn = this.userService.isLoggedIn()
  }


  updateLoginStatus() {
    this._isLoggedIn = this.userService.isLoggedIn();
    this.id = this.userService.get_user_id();
    if (this._isLoggedIn) {
      console.log("Пользователь авторизован:", this.id);
    } else {
      console.log("Пользователь не авторизован");
    }
  }

  logout() {
    this.userService.logout_on_site();
    this.updateLoginStatus();
  }

  showLoginModal() {
    this.showLogin = true;
  }

  closeLoginModal() {
    this.showLogin = false;
  }

  showRegisterModal() {
    this.showRegister = true;
  }

  closeRegisterModal() {
    this.showRegister = false;
  }

  navigate() {
    this.router.navigate(['/user-profile', this.id]);
  }
}