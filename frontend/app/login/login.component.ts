import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators, AbstractControl, FormControl } from '@angular/forms';
import Validation from '../utils/validation';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { AuthService } from '../auth.service';
import { User } from '../user';
import { UserService } from '../user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  @Output() loginComplete = new EventEmitter<void>();

  form: FormGroup = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
  });
  submitted = true;
  is_visible = true;
  user_id : string = "";
  username : string = "";

  hideComponent(){
    this.is_visible = false;
  }

  viewComponent(){
    this.is_visible = true;
  }

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router, private userService: UserService) {}


  ngOnInit() {
    this.form = this.formBuilder.group(
      {
        username: [
          '',
          [
            Validators.required,
            Validators.minLength(6),
            Validators.maxLength(20)
          ]
        ],
        password: [
          '',
          [
            Validators.required,
            Validators.minLength(6),
            Validators.maxLength(40)
          ]
        ],
      },
    );
    const placeholder = {
      username: "user1111",
      password: "password1",
  };
  this.form.patchValue(placeholder);
  }

  get f(): { [key: string]: AbstractControl } {
    return this.form.controls;
  }

  onSubmit() {
    this.submitted = true;

    if (this.form.invalid) {
      return;
    }

    this.username = this.form.value.username;
    this.userService.login(this.form.value).subscribe(
      response => {
        this.userService.getUserbyUsername(this.username).subscribe({
          next: (data: any) => { 
            this.user_id = data;
            this.loginComplete.emit();
            this.userService.login_on_site(this.user_id);
            this.router.navigate(['/user-profile', this.user_id]);
        }, error: (err: any) => console.error('Error with get Username', err),
      });
      console.log("Login successful", response);
    }, error => {
      console.log("Login failed", error)
    }
    );
  }

  onReset(): void {
    this.submitted = false;
    this.form.reset();
  }
}