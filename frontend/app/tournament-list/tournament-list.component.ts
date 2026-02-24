import { Component } from '@angular/core';
import { Tournament } from '../tournament';
import { TournamentService } from '../tournament.service'

@Component({
  selector: 'app-tournament-list',
  templateUrl: './tournament-list.component.html',
  styleUrl: './tournament-list.component.css'
})
export class TournamentListComponent {
  tournaments: Tournament[] = [];

    constructor(private TournamentService: TournamentService) {
      this.tournaments = new Array<Tournament>();
  }
    ngOnInit() {
      this.TournamentService.getTournaments().subscribe({next:(data: Tournament[]) => this.tournaments=data});
    }
}
