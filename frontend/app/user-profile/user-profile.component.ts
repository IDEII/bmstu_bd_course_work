import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { User } from '../user'
import { UserService } from '../user.service';
import { FormGroup, FormControl, FormBuilder } from '@angular/forms';
import { FighterComponent } from '../fighter/fighter.component';
import { ClubProfileComponent } from '../club-profile/club-profile.component';
import { OrganazerProfileComponent } from '../organazer-profile/organazer-profile.component';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {
  user: User = {
    user_profile_id: "",
    username: "",
    profile_type: 0,
    e_mail: "",
    linked_id_1: "",
    linked_id_2: "",
    linked_id_3: "",
    password: ""
};

user_id: string = "";
isEditing: boolean = false;
showConfirmDelete: boolean = false;
chngPassword: boolean = false;
form: FormGroup;

fighterId : string | null = "";
organazerId : string | null = "";
clubId : string | null = "";

fighterProfile = FighterComponent;
clubProfile = ClubProfileComponent;
organazerProfile = OrganazerProfileComponent;

constructor(
  private router: Router,
  private formBuilder: FormBuilder,
  private route: ActivatedRoute,
  private userService: UserService
) {
  this.form = this.formBuilder.group({
    user_profile_id: new FormControl(''),
    username: new FormControl(''),
    profile_type: new FormControl(0),
    e_mail: new FormControl(''),
    linked_id_1: new FormControl(''),
    linked_id_2: new FormControl(''),
    linked_id_3: new FormControl(''),
    password: new FormControl(''),
  });
}

  async ngOnInit() {
  const id = this.route.snapshot.paramMap.get('id');
  if (id) {
    this.user_id = id;
    await this.loadUser(this.user_id);
  }
  this.clubId = this.user.linked_id_2;
  this.fighterId = this.user.linked_id_3;
  this.organazerId = this.user.linked_id_1;
  console.log("user", this.user);
}

async loadUser(id: string) {
  try {
  const data = await this.userService.getUser(id).toPromise();
  this.user = data;
  this.form.patchValue(this.user); 
  console.log('Данные пользователя загружены:', this.user); 
} catch (error) {
  console.error('Ошибка при загрузке пользователя:', error);
}
}

changePassword() {
  this.chngPassword = !this.chngPassword
}

editUser() {
  if (!this.form.valid) {
    return;
  }
  if (!this.chngPassword) {
    const password = {
      password: ""
    }
    this.form.patchValue(password)
  }
  this.userService.updateUser(this.user_id, this.form.value).subscribe(
    response => {
      console.log('Профиль успешно обновлен');
      this.isEditing = false;
      this.loadUser(this.user_id); 
    },
    error => {
      console.error('Ошибка при обновлении профиля', error);
    }
  );
}

cancelEdit() {
  this.isEditing = false;
  this.form.reset();
  this.form.patchValue(this.user); 
}

deleteUser() {
  this.userService.logout_on_site();
  this.userService.deleteUser(this.user.user_profile_id).subscribe(() => {
    console.log('Пользователь удален');
    this.router.navigate(['/home']);
  }, error => {
    console.error('Ошибка при удалении пользователя', error);
  });
}
}