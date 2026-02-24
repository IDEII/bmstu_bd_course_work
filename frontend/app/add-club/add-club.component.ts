import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators, AbstractControl, FormControl, FormsModule } from '@angular/forms';
import { Club } from '../club';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import Validation from '../utils/validation';
import { ClubService } from '../club.service';
import { UserService } from '../user.service';

@Component({
  selector: 'app-add-club',
  templateUrl: './add-club.component.html',
  styleUrl: './add-club.component.css'
})

export class AddClubComponent implements OnInit {
  club : Club = {
    user_id: '',
    id: '0', 
    title: '', 
    description: '',
    address: '',
    contact: '',
    rating: '',
    founded_date: new Date(),
  };
  user_id : string = "";
  form: FormGroup = new FormGroup ({
    title: new FormControl(''),
    description: new FormControl(''),
    address: new FormControl(''),
    contact: new FormControl(''),
    rating: new FormControl(''),
    founded_date: new FormControl(new Date()),
  })
  constructor(private formBuilder: FormBuilder, private clubService: ClubService, private router: Router, private userService: UserService) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      title: ["",

      ],
      description: ["",

      ],
      
      address: ["",

      ],
      
      contact: ["",

      ],
      rating: ["",

      ],
      founded_date: [new Date(),

      ],
    },
    );
    this.user_id = this.userService.get_user_id();

  }
  addClub(): void {
    this.clubService.addClub(this.form.value, this.user_id).subscribe(
      response => {
        console.log('good thing', response);
        this.router.navigate(['/clubs']);
    }, error => {
      console.error('bad thing', error);
    }
  );
  }
}
