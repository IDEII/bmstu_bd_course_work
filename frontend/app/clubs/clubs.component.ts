import {Component, OnInit, TemplateRef, ViewChild} from "@angular/core";
import {NgTemplateOutlet} from "@angular/common";
import { FormsModule }   from "@angular/forms";
import { HttpClientModule }   from "@angular/common/http";
import { Club } from "../club";
import { ClubMember } from '../club-member'
import { ClubService } from "../club.service";
import { Router } from '@angular/router'

@Component({
  selector: 'app-clubs',
  templateUrl: './clubs.component.html',
  styleUrl: './clubs.component.css'
})


export class ClubsComponent implements OnInit {
    clubs: Club[] = [];
    statusMessage: string = '';

    constructor(private clubService: ClubService) {
      
  }
    ngOnInit() {
        this.clubService.getClubs().subscribe({next:(data: Club[]) => this.clubs=data});
    }

}