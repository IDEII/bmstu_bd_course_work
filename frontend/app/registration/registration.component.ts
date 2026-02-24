import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators, AbstractControl, FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { AuthService } from '../auth.service';
import { UserService } from '../user.service';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {
  @Output() registrationComplete = new EventEmitter<void>();
  registrationForm: FormGroup = new FormGroup ({
    user_profile_id: new FormControl(''),
    username: new FormControl(''),
    profile_type: new FormControl(0),
    e_mail: new FormControl(''),
    linked_id_1: new FormControl(''),
    linked_id_2: new FormControl(''),
    linked_id_3: new FormControl(''),
    password: new FormControl(''),
    confirmPassword: new FormControl(''),
  });
  submitted = true;
  user_id : string =  "";
  e_mail : string = "";
  is_visible = true;

  hideComponent(){
    this.is_visible = false;
  }

  viewComponent(){
    this.is_visible = true;
  }

  constructor(private formBuilder: FormBuilder, private http: HttpClient, private router: Router, private userService: UserService) {
    
    this.registrationForm = this.formBuilder.group(
      {
        username: ['', [Validators.required, Validators.minLength(6)]],
        profile_type: [0, ],
        e_mail: ['', [Validators.required, Validators.email]],
        linked_id_1: ['',],
        linked_id_2: ['',],
        linked_id_3: ['',],
        password: ['', [Validators.required, Validators.minLength(6)]],
        confirmPassword: ['', [Validators.required]],
      },
      {
        validator: this.passwordMatchValidator
      }
    );
  }

  passwordMatchValidator(form: FormGroup) {
    return form.get('password')?.value === form.get('confirmPassword')?.value ? null : { mismatch: true };
  }

  ngOnInit() {
    const placeholder = {
        username: "testUser",
        e_mail: "testEmail@email.com",
        password: "1234567890",
        confirmPassword: "1234567890",
    };
    this.registrationForm.patchValue(placeholder);
  }


  async onSubmit() {
    this.submitted = true;
  
    if (this.registrationForm.invalid) {
      return;
    }
  
    const username = this.registrationForm.value.username;
  
    this.registrationForm.removeControl('confirmPassword');
  
    console.log("user", username);
  
    try {
      const registrationResponse = await this.userService.register(this.registrationForm.value).toPromise();
      console.log('Registration successful', registrationResponse);
  
      const data = await this.userService.getUserbyUsername(username).toPromise();
      if (data) {
        this.user_id = data;
        this.registrationComplete.emit();
      } else {
        return;
      }
  
      this.userService.login_on_site(this.user_id);
      console.log(this.user_id);
      this.router.navigate(['/user-profile', this.user_id]);
      
    } catch (error) {
      console.error('Registration failed', error);
      this.addConfirmPasswordControl();
    }
  
    console.log(username);
  }

  addConfirmPasswordControl() {
    this.registrationForm.addControl('confirmPassword', this.formBuilder.control('', [Validators.required]));
  }

  onReset(): void {
    this.submitted = false;
    this.registrationForm.reset();
  }
}