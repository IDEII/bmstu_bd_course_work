import { Component } from '@angular/core';
import { Fighter } from '../fighter';
import { FighterService } from '../fighter.service';
import { Router } from '@angular/router';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';

@Component({
  selector: 'app-add-fighter',
  templateUrl: './add-fighter.component.html',
  styleUrl: './add-fighter.component.css'
})
export class AddFighterComponent {
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
  form: FormGroup = new FormGroup ({
    name: new FormControl(''),
    description: new FormControl(''),
    country: new FormControl(''),
    rating: new FormControl(''),
    category: new FormControl(0),
    birthday: new FormControl(new Date()),
  })
  user_id: string = "";
  
  constructor(private formBuilder: FormBuilder,private fighterService: FighterService, private router: Router, private userService: UserService) { }
  ngOnInit(): void {
    this.form = this.formBuilder.group({
      name: ["",

      ],
      description: ["",

      ],
      country: ["",

      ],
      rating: ["",

      ],
      category: [0,

      ],
      birthday: [new Date(),

      ],
    },
    );
    this.user_id = this.userService.get_user_id();

  }
  addFighter(): void {
  this.fighterService.addFighter(this.form.value, this.user_id).subscribe(
    response => {
      console.log('good thing', response);
      this.router.navigate(['/fighters']);
  }, error => {
    console.error('bad thing', error);
  }
);
}
}
