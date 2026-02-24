import { Component, Input } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Fighter } from '../fighter';
import { FighterService } from '../fighter.service';
import { FormGroup, FormControl, FormBuilder } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { RequestOnTournamentComponent } from '../request-on-tournament/request-on-tournament.component';
import { UserService } from '../user.service';

@Component({
  selector: 'app-fighter',
  templateUrl: './fighter.component.html',
  styleUrl: './fighter.component.css'
})
export class FighterComponent {
  @Input() fighterID : string = "";
  editable : boolean = false; 
  fighter : Fighter = {
         user_id: '0',
         club_id: '0',
         id: '0',
         name: "",
         description: "",
         country: "",
         birthday: new Date(),
         rating: "",
         category: 0,
}
app_request = RequestOnTournamentComponent;
isEditing: boolean = false;
showConfirmDelete: boolean = false; 
form: FormGroup = new FormGroup ({
  club_id: new FormControl(''),
  id: new FormControl(''),
  name: new FormControl(''),
  description: new FormControl(''),
  country: new FormControl(''),
  birthday: new FormControl(new Date()),
  rating: new FormControl(''),
  category: new FormControl(0),

})
user_id : string = "";

constructor(private datepipe: DatePipe, private formBuilder: FormBuilder, private fighterService: FighterService, private route: ActivatedRoute, private router: Router, private userService: UserService) {
}
ngOnInit() {
  this.user_id = this.userService.get_user_id();

  this.form = this.formBuilder.group({
    club_id: ["",],
    id: ["",],
    name: ["",],
    description: ["",],
    country: ["",],
    birthday: [new Date(),],
    rating: ["",],
    category: [0,],
  })
  const placeholder = {
    name: "1",
    description: "1",
    country: "1",
    rating: "1",
    category: 0,
  } 
  this.form.patchValue(placeholder)
  var id : string | null = "";
  if (this.fighterID !== "") {
      id = this.fighterID;
      this.editable = true;
  } else {
    id = this.route.snapshot.paramMap.get('id');  }
  if (id !== null) {
  this.fighterService.getFighter(id).subscribe(data => {
    this.fighter = data;
  });
}
}

delFighter() {
  this.fighterService.deleteFighter(this.fighter.id).subscribe(() => {
    console.log('боец успешно удален');
    this.router.navigate(['/fighters']);
  }, error => {
    console.error('Ошибка при удалении бойца', error);
  });
}
updFighter () {
  const tempDate = this.datepipe.transform(this.fighter.birthday, 'yyyy-MM-dd');
  const formData = {
    id: this.fighter.id,
    club_id: this.fighter.club_id,
    birthday: tempDate
  }
  this.form.patchValue(formData)
  console.log(this.form.value)
  const id = this.fighter.id
  if (id !== null) {
    this.fighterService.updateFighter(id, this.form.value).subscribe(
      response => {
        console.log('good thing');
        this.isEditing = false;
    }, error => {
      console.log('bad thing', error)
    }
    );
    this.fighterService.getFighter(id).subscribe(data => {
      this.fighter = data;  });

  }
}

cancelEdit() {
  this.isEditing = false;
  this.form = this.formBuilder.group({
    club_id: ["",],
    id: ["",],
    name: ["",],
    description: ["",],
    country: ["",],
    birthday: [new Date(),],
    rating: ["",],
    category: [0,],
  })
}
}
