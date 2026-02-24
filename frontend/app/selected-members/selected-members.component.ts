import { Component, Input, OnInit } from '@angular/core';
import { TournamentService } from '../tournament.service'
import { Participant } from '../tournament';
import { ActivatedRoute } from '@angular/router';


@Component({
  selector: 'app-selected-members',
  templateUrl: './selected-members.component.html',
  styleUrls: ['./selected-members.component.css']
})

export class SelectedMembersComponent {
  
}