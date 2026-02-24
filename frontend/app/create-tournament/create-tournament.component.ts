import { Component, Input, OnInit} from '@angular/core';
import { FormBuilder, FormGroup, Validators, AbstractControl, FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { Tournament } from '../tournament'
import { TournamentService } from '../tournament.service'
import { UserService } from '../user.service';
@Component({
  selector: 'app-create-tournament',
  templateUrl: './create-tournament.component.html',
  styleUrl: './create-tournament.component.css'
})
export class CreateTournamentComponent {
  @Input() organazer_id = "";
  tournament : Tournament = {
    id: '',
    name: '',
    address: '',
    startDate: new Date(),
    endDate: new Date(),
    organazer: '',
    rounds: '',
    category: 0,
  };
  form: FormGroup = new FormGroup ({
    id: new FormControl(''),
    name: new FormControl(''),
    address: new FormControl(''),
    organazer: new FormControl(this.organazer_id),
    startDate: new FormControl(new Date()),
    endDate: new FormControl(new Date()),
    rounds: new FormControl(''),
    category: new FormControl(''),
  })
  
  constructor(private formBuilder: FormBuilder, private tournamentService: TournamentService, private router: Router, private userService: UserService) { }
  ngOnInit(): void {
    this.form = this.formBuilder.group({
      id: ['',],
      name: ['', Validators.required],
      address: ['',],
      startDate: [new Date(), Validators.required],
      endDate: [new Date(), Validators.required],
      organazer: ['',],
      rounds: ['', Validators.required],
      category: [0, ]
    });
  }
  
  addTournament(): void {
    const formData = {
      organazer: this.organazer_id
    }
    this.form.patchValue(formData)
    this.tournamentService.addTournament(this.form.value).subscribe(response => {
      console.log('good thing', response);
      this.router.navigate(['/tournaments']);
  }, error => {
    console.error('bad thing', error);
  }
);
  }
}