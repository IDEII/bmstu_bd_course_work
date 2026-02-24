import { Component, OnInit } from '@angular/core';
import { OrganazerService } from '../organazer.service'
import { Organazer } from '../organazer';
import { Router } from '@angular/router';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { UserService } from '../user.service';

@Component({
  selector: 'app-add-organazer',
  templateUrl: './add-organazer.component.html',
  styleUrl: './add-organazer.component.css'
})
export class AddOrganazerComponent implements OnInit {
  organazer : Organazer = {
    user_id: "",
     id: "",
     title: "",
     description: "",
     address: "",
     contact: "",
     founded_date: new Date()
  };
  form: FormGroup = new FormGroup({
    title: new FormControl(''),
    description: new FormControl(''),
    address: new FormControl(''),
    contact: new FormControl(''),
    founded_date: new FormControl(new Date()),
  })
  user_id: string = "";

  constructor(private formBuilder: FormBuilder, private organazerService: OrganazerService, private router: Router, private userService: UserService) { }
  ngOnInit(): void {
    this.form = this.formBuilder.group({
      title: ['', Validators.required],
      description: ['', Validators.required],
      address: ['', Validators.required],
      contact: ['', Validators.required],
      founded_date: [new Date(), Validators.required]
    });
    this.user_id = this.userService.get_user_id();

  }
  addOrganazer(): void {
    this.organazerService.addOrganazer(this.form.value, this.user_id).subscribe(
      response => {
        console.log('good thing', response);
        this.router.navigate(['/organazers']);
    }, error => {
      console.error('bad thing', error);
    }
  );
  
}
}